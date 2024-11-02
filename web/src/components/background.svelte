<script lang="ts">
	import type { Snippet } from "svelte";

    const {children} : {children: Snippet} = $props();
</script>

<div class="h-full w-full bg-black">
 {@render children?.()}
</div>