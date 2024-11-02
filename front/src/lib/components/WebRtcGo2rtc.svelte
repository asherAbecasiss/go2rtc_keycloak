<script>
  import { onMount, onDestroy } from 'svelte'
  import { GetPublicIp, Getgo2rtcBasePath } from '$lib/utils/utils.js'
  import { keycloak } from '$lib/keycloak.js'

  export let suuid

  const apiUrl = GetPublicIp()
  let token = keycloak.token
  let stream = new MediaStream()
  const pc = new RTCPeerConnection({
    iceServers: [{ urls: 'stun:stun.l.google.com:19302' }],
  })

  async function PeerConnection(media) {
    const localTracks = []

    if (/camera|microphone/.test(media)) {
      const tracks = await getMediaTracks('user', {
        video: media.indexOf('camera') >= 0,
        audio: media.indexOf('microphone') >= 0,
      })
      tracks.forEach((track) => {
        pc.addTransceiver(track, { direction: 'sendonly' })
        if (track.kind === 'video') localTracks.push(track)
      })
    }

    if (media.indexOf('display') >= 0) {
      const tracks = await getMediaTracks('display', {
        video: true,
        audio: media.indexOf('speaker') >= 0,
      })
      tracks.forEach((track) => {
        pc.addTransceiver(track, { direction: 'sendonly' })
        if (track.kind === 'video') localTracks.push(track)
      })
    }

    if (/video|audio/.test(media)) {
      const tracks = ['video', 'audio']
        .filter((kind) => media.indexOf(kind) >= 0)
        .map(
          (kind) =>
            pc.addTransceiver(kind, { direction: 'recvonly' }).receiver.track,
        )
      localTracks.push(...tracks)
    }

    stream.srcObject = new MediaStream(localTracks)

    return pc
  }

  async function getMediaTracks(media, constraints) {
    try {
      const stream =
        media === 'user'
          ? await navigator.mediaDevices.getUserMedia(constraints)
          : await navigator.mediaDevices.getDisplayMedia(constraints)
      return stream.getTracks()
    } catch (e) {
      console.warn(e)
      return []
    }
  }

  async function connect(media) {
    const pc = await PeerConnection(media)

    const url = new URL(
      Getgo2rtcBasePath() + 'api/ws' + location.search,
      'http://' + location.host,
    )
    // const url ="http://localhost:1985"
    url.token = 'op'
    url.port = '1985'
    url.search = '?src=' + suuid + '&media=video+audio&token=' + token

    const ws = new WebSocket('ws' + url.toString().substring(4))

    ws.addEventListener('open', () => {
      pc.addEventListener('icecandidate', (ev) => {
        if (!ev.candidate) return
        const msg = { type: 'webrtc/candidate', value: ev.candidate.candidate }
        ws.send(JSON.stringify(msg))
      })

      pc.createOffer()
        .then((offer) => pc.setLocalDescription(offer))
        .then(() => {
          const msg = { type: 'webrtc/offer', value: pc.localDescription.sdp }
          ws.send(JSON.stringify(msg))
        })
    })

    ws.addEventListener('message', (ev) => {
      const msg = JSON.parse(ev.data)
      if (msg.type === 'webrtc/candidate') {
        pc.addIceCandidate({ candidate: msg.value, sdpMid: '0' })
      } else if (msg.type === 'webrtc/answer') {
        pc.setRemoteDescription({ type: 'answer', sdp: msg.value })
      }
    })
  }
  onMount(async () => {
    const media = new URLSearchParams(location.search).get('media')
    await connect(media || 'video+audio')
  })

  onDestroy(() => {
    console.log('Date Component removed webRtc')
    pc.close()
  })
</script>


<video
  bind:this={stream}
  height="240px"
  width="320px"
  class="video w-full aspect-video rounded-lg "
  autoplay
  controls
  playsinline
  muted />
