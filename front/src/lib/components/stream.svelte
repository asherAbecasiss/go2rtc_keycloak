<script>
  import {
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
    Button,
  } from 'flowbite-svelte'
  import { onMount } from 'svelte'
  import { GetPublicIp, Getgo2rtcBasePath } from '$lib/utils/utils.js'
  import ImageSync from './ImageSync.svelte'
  import { goto } from '$app/navigation'
  export let stream
  const apiUrl = GetPublicIp()
  const base_path=Getgo2rtcBasePath()
  function goTo(path) {
    goto(path)
  }
</script>

<Table>
  <TableHead>
    <TableHeadCell>Image</TableHeadCell>
    <TableHeadCell>camera name</TableHeadCell>
    <TableHeadCell>play WEBRTC</TableHeadCell>
    <TableHeadCell>play HLS</TableHeadCell>
  </TableHead>
  <TableBody tableBodyClass="divide-y">

    {#each stream as s}
      <TableBodyRow>

        <TableBodyCell>
          <ImageSync
            url={`${apiUrl}${base_path}api/frame.jpeg?src=` + s.name}
            size={'w-18 h-16 '} />
        </TableBodyCell>
        <TableBodyCell>{s.name}</TableBodyCell>
        <TableBodyCell>
          <Button on:click={goTo("/webrtc/"+s.name)} outline>WebRTC</Button>
        </TableBodyCell>
        <TableBodyCell>
          <Button on:click={goTo("/hls/"+s.name)} outline>HLS</Button>
        </TableBodyCell>

      </TableBodyRow>
    {/each}

  </TableBody>
</Table>
