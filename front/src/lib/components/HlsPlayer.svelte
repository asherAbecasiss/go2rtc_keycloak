<script>
    import Hls from 'hls.js'
    import { onMount, onDestroy } from 'svelte'
    import { GetPublicIp } from '$lib/utils/utils.js'
     import { keycloak } from '$lib/keycloak.js'
    const apiUrl = GetPublicIp()
    export let suuid
    let videoElement
    let hls
    let urlLive = apiUrl + '/api/stream.m3u8?src=' + suuid + '&mp4'
    console.log(urlLive)
  
    onMount(() => {
      hls = new Hls({
        xhrSetup: (xhr) => {
          xhr.setRequestHeader('Authorization', keycloak.token)
        },
      })
  
      hls.on(Hls.Events.MANIFEST_PARSED, function () {
        console.log('HLS manifest parsed', hls.levels)
      })
  
      hls.loadSource(urlLive)
      hls.attachMedia(videoElement)
    })
  
    onDestroy(() => {
      console.log('Date Component removed HLS feed')
      hls.destroy()
      hls = null
    })
  </script>
  
  <video
    id="asher"
    class=" rounded-lg"
    width="100%"
    height="298"
    bind:this={videoElement}
    autoplay="true"
    controls>
    <track kind="captions" />
  </video>
  