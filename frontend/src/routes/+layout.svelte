<script lang="ts">
	import './layout.css';
	import { setContext } from 'svelte';
	import { createAuth } from '$lib/contexts/auth.svelte';

	let { children } = $props();

	const auth = createAuth();
	setContext('auth', auth);
</script>

<svelte:head><link rel="icon"/></svelte:head>
{@render children()}
