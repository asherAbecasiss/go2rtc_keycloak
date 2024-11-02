<script>
  import '../app.css'

  import { onMount } from 'svelte'
  import { initializeKeycloak, keycloak, getUserInfo } from '$lib/keycloak.js'
  import { page } from '$app/stores'
  import { Navbar, NavBrand, NavLi, NavUl, NavHamburger,Button } from 'flowbite-svelte'
  $: activeUrl = $page.url.pathname
  let activeClass =
    'text-white bg-green-700 md:bg-transparent md:text-green-700 md:dark:text-white dark:bg-green-600 md:dark:bg-transparent'
  let nonActiveClass =
    'text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-green-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent'

  let authenticated = false
  let user_email = ''

  onMount(async () => {
    authenticated = await initializeKeycloak()
    if (authenticated) {
      user_email = getUserInfo()
    }
  })

  function logout() {
    keycloak.logout()
  }
</script>

{#if authenticated}
  <!-- Show the main app content -->
  <Navbar>
    <NavBrand href="/">

      <span
        class="self-center whitespace-nowrap text-xl font-semibold
        dark:text-white">
        go2keycloak
      </span>
    </NavBrand>
    <NavHamburger />
    <NavUl {activeUrl} {activeClass} {nonActiveClass}>
      <NavLi href="/">Home</NavLi>
      <NavLi href="/all">All cameras</NavLi>

      <NavLi href="#" on:click={logout} >Logout {user_email}</NavLi>
    </NavUl>
  </Navbar>
  <slot />

{:else}
  <!-- Loading or redirect message -->
  <p>Loading...</p>
{/if}
