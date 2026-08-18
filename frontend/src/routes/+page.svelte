<script lang="ts">
  let username = $state('');
  let password = $state('');

  function handleLogin(event: SubmitEvent) {
    event.preventDefault();
    // Handle login logic here
  }
</script>

<main class="flex min-h-screen items-center justify-center bg-white text-neutral-900 font-sans text-sm">
  <div class="w-full max-w-sm border border-neutral-200 p-8 bg-white shadow-sm">
    <div class="mb-8">
      <h1 class="text-base font-semibold tracking-tight m-0 mb-1 text-blue-950">IPS Admin</h1>
      <p class="text-xs text-neutral-500 m-0">Enter your credentials to access the system.</p>
    </div>

    <form onsubmit={handleLogin} class="space-y-4">
      <div>
        <label class="block text-xs font-medium uppercase tracking-wider text-neutral-500 mb-1.5" for="username">
          Username
        </label>
        <input
          id="username"
          type="text"
          bind:value={username}
          required
          class="w-full border border-neutral-300 p-2 text-sm bg-white focus:outline-none focus:border-blue-950"
        />
      </div>

      <div>
        <label class="block text-xs font-medium uppercase tracking-wider text-neutral-500 mb-1.5" for="password">
          Password
        </label>
        <input
          id="password"
          type="password"
          bind:value={password}
          required
          class="w-full border border-neutral-300 p-2 text-sm bg-white focus:outline-none focus:border-blue-950"
        />
      </div>

      <button
        type="submit"
        class="w-full bg-blue-950 text-white font-medium py-2.5 px-4 text-xs uppercase tracking-wider hover:bg-blue-900 transition-colors mt-2"
      >
        Sign In
      </button>
    </form>
  </div>
</main>
