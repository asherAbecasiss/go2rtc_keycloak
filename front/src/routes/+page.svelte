<script>
  import Stream from '$lib/components/stream.svelte'
  import { keycloak } from '$lib/keycloak.js'
  import { onMount } from 'svelte'
  import { GetPublicIp, Getgo2rtcBasePath } from '$lib/utils/utils.js'
  let Data = null // Store fetched data here
  let error = null // Store any errors
  let stream = []
  let loading = true // Loading state to show while waiting for data

  const apiUrl = GetPublicIp()
  const base_path=Getgo2rtcBasePath()
  console.log(keycloak.token);
  onMount(async () => {
    try {
      const response = await fetch(`${apiUrl}${base_path}/api/streams`, {
        method: "GET",
        headers: {
          Authorization: `${keycloak.token}`,
          "Content-Type": "application/json",
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
</script>

{#if loading}
  <p>Loading...</p>
{:else if error}
  <p>Error: {error}</p>
{:else}
  <Stream {stream} />
{/if}
