<script lang="ts">
    import { onMount, onDestroy } from 'svelte';

    let currentTime = $state(new Date());
    let intervalId: NodeJS.Timeout;

    onMount(() => {
        intervalId = setInterval(() => {
            currentTime = new Date();
        }, 200);

        return () => {
            clearInterval(intervalId);
        };
    });
    const twelveHour = (t: Date) => {
        let h = t.getHours();
        return (h % 12 || 12).toString();
    };

    let hours: string = $derived(twelveHour(currentTime));
    let minutes = $derived(currentTime.getMinutes().toString().padStart(2, '0'));
    let seconds = $derived(currentTime.getSeconds().toString().padStart(2, '0'));
    let ampm = $derived(currentTime.getHours() >= 12 ? 'pm' : 'am');

</script>

<div class="flex flex-row text-white h-[50%] bg-gray-700">
    <div>it is</div>
    <div>{hours}:{minutes}:{seconds} {ampm}</div>
</div>