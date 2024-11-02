<script>
    export let data

    import HlsPlayer from '$lib/components/HlsPlayer.svelte';
      import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte'
  </script>
  
  
  <div class="flex-1 flex ">
  
    <div class="flex flex-col w-full example">
      <div class="flex-col w-full pt-2 px-2 ">
        <Breadcrumb aria-label="Default breadcrumb ">
          <BreadcrumbItem
            homeClass="inline-flex items-center text-sm font-medium
            hover:text-gray-400 dark:text-gray-400 dark:hover:text-white
            text-black"
            href="/"
            home>
            Home
          </BreadcrumbItem>
  
  
          <BreadcrumbItem
            spanClass="ms-1 text-sm font-medium text-orange-600 md:ms-2
            dark:text-gray-400">
            HLS {data.cameraName}
          </BreadcrumbItem>
        </Breadcrumb>
      </div>
  
      <div class=" grid grid-cols-6 gap-4  rounded-t-lg ">
  
        <!-- <div>
          <MapModel suuid={data.cameraName} s={'xs:invisible lg:visible  lg:h-full  p-2  border border-gray-600 rounded-b-lg  '} />
        </div> -->
  
        <div class="col-start-2 col-span-4 relative bg-blend">
          <div class="">
            <div class="relative flex justify-center overflow-hidden m-2">
              <div class="">
                <div
                  class="card bg-black rounded-lg p-1 r "
                  style={`--color: 255, 93, 93;"`}>
                  <div class="flex flex-row ml-2 mt-1 mb-1 items-center gap-1 ">
  
                   
  
                  </div>
  
                  <HlsPlayer suuid={data.cameraName} />
  
                </div>
              </div>
  
            </div>
          </div>
        </div>
  
      </div>
    </div>
  </div>
  