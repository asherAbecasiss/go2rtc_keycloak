<script>
  import Stream from '$lib/components/stream.svelte'
  import { keycloak } from '$lib/keycloak.js'
  import { onMount } from 'svelte'
  import { gridSize, player_t } from '$lib/store'
  import HlsPlayer from '$lib/components/HlsPlayer.svelte'
  import {
    Navbar,
    NavBrand,
    NavLi,
    NavUl,
    NavHamburger,
    Dropdown,
    DropdownItem,
    Button,
    Input,
  } from 'flowbite-svelte'
  import { ChevronDownOutline } from 'flowbite-svelte-icons'
  import WebRtcGo2rtc from '$lib/components/WebRtcGo2rtc.svelte'
  import { GetPublicIp, Getgo2rtcBasePath } from '$lib/utils/utils.js'
  const apiUrl = GetPublicIp()
  const base_path = Getgo2rtcBasePath()
  let Data = null // Store fetched data here
  let error = null // Store any errors
  let stream = []
  let loading = true // Loading state to show while waiting for data
  let pl_t

  player_t.subscribe((value) => (pl_t = value))
  let grid_size
  gridSize.subscribe((value) => (grid_size = value))
  $: gSize = $gridSize

  onMount(async () => {
    try {
      const response = await fetch(`${apiUrl}${base_path}/api/streams`, {
        headers: {
          Authorization: keycloak.token,
          'Content-Type': 'application/json',
        },
      })

      if (!response.ok) {
        throw new Error('Failed to fetch data')
      }

      // Parse JSON response
      Data = await response.json()

      for (const [key, value] of Object.entries(Data)) {
        stream.push({ name: key })
      }
      console.log(stream)
    } catch (err) {
      error = err.message
    } finally {
      loading = false // Set loading to false after completion
    }
  })

  const playrType = (j) => () => {
    player_t.set(j)
  }
  const handleClick = (j) => () => {
    gridSize.set(j)
  }
</script>

<svelte:head>
  <title>LIVE</title>
</svelte:head>

{#if loading}
  <p>Loading...</p>
{:else if error}
  <p>Error: {error}</p>
{:else}
  <!-- <p class="bar "> -->

  <div class="flex justify-between bg-black p-1">
    <div class="order-last">

      <Button size="xs">
        Player
        <ChevronDownOutline class="w-6 h-6 text-white dark:text-white" />
      </Button>
      <Dropdown class="w-44 z-20">
        <DropdownItem on:click={playrType('hls')}>HLS</DropdownItem>
        <DropdownItem on:click={playrType('webrtc')}>WebRTC</DropdownItem>

      </Dropdown>

      <Button size="xs">
        Grid Size
        <ChevronDownOutline class="w-6 h-6 ms-2 text-white dark:text-white" />
      </Button>

      <Dropdown class="w-44 z-20">
        <DropdownItem
          on:click={handleClick('grid xl:grid-cols-2 md:grid-cols-1')}>
          2x2
        </DropdownItem>
        <DropdownItem
          on:click={handleClick('grid xl:grid-cols-3 md:grid-cols-1')}>
          3x3
        </DropdownItem>
        <DropdownItem
          on:click={handleClick('grid xl:grid-cols-4 md:grid-cols-1')}>
          4x4
        </DropdownItem>
        <DropdownItem
          on:click={handleClick('grid xl:grid-cols-5 md:grid-cols-1')}>
          5x5
        </DropdownItem>
        <DropdownItem
          on:click={handleClick('grid xl:grid-cols-6 md:grid-cols-1')}>
          6x6
        </DropdownItem>
        <DropdownItem
          on:click={handleClick('grid xl:grid-cols-8 md:grid-cols-1')}>
          8x8
        </DropdownItem>
        <DropdownItem
          on:click={handleClick('grid xl:grid-cols-10 md:grid-cols-1')}>
          10x10
        </DropdownItem>
      </Dropdown>

    </div>

    <div class="text-white font-bold">
      <span class="live-icon" />
      {#if pl_t === 'webrtc'}RTC{:else}HLS{/if}

    </div>
  </div>

  <div class="grid grid-cols-8 ">
    <div class="col-span-8 ">
      <main class=" z-0">
        <div class="{grid_size} ">
          {#each stream as s (s)}
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <div
              class="
              "
              id={s}
              animate:flip={{ duration: dragDuration }}
              draggable="true"
              on:dragstart={() => (draggingCard = s)}
              on:dragend={() => (draggingCard = undefined)}
              on:dragenter={() => swapWith(s)}
              on:dragover|preventDefault>
              <div class="relative bg-blend ">
                <div class="">
                  <div
                    class="relative flex justify-center overflow-hidden m-2 r
                    bg-black rounded-lg "
                    id={s.name}>
                    <div class="">
                      <div class=" p-1 " style={`--color: 255, 93, 93;"`}>
                        <div
                          class="flex flex-row ml-2 mt-1 mb-1 items-center gap-1
                          ">

                          <span class="text-red-600">{pl_t}</span>

                          <span id="sep" class="text-gray-300">•</span>

                          <p class=" text-gray-200 rounded-lg">

                            <span>{s.name}</span>
                          </p>

                        </div>
                        {#if pl_t === 'webrtc'}
                          <WebRtcGo2rtc suuid={s.name} />
                        {:else}
                          <HlsPlayer suuid={s.name} />
                        {/if}
                      </div>
                    </div>

                  </div>
                </div>
              </div>
            </div>
          {/each}
        </div>
      </main>
    </div>

  </div>
{/if}
