<script lang="ts">
    import { onMount } from "svelte";
    import OutputIr from "./output-ir.svelte";
    import OutputRan from "./output-ran.svelte";
    import OutputAfterImports from "./output-after-imports.svelte";
    import OutputWasm from "./output-wasm.svelte";

    export let response_text : string;
    export let response_ir : string;
    export let response_after_import: string;
    export let stdout_compiler: string;
    export let response_wasm: string;
    export let tabs: any[] = [
        {
            label: "Output",
            value: 1,
            component: OutputRan,
        }, 
        {
            label: "IR",
            value: 2,
            component: OutputIr
        },
        {
            label: "After Imports",
            value: 3,
            component: OutputAfterImports
        },
        {
            label: "WASM",
            value: 4,
            component: OutputWasm
        }
    ];
    export let activeTabValue: number = 1;

    type Tab = {
        value: number,
        label: string,
       component: OutputIr | OutputRan 
    };
    onMount(() => {
    // Set default tab value
    if (Array.isArray(tabs) && tabs.length && tabs[0].value) {
      activeTabValue = tabs[0].value;
    }
    });
    function handleClick(tabValue: number){
        activeTabValue = tabValue;
        return null;
    }
</script>

<ul>
    {#each tabs as tab}
    <li class={activeTabValue === tab.value ? 'active' : ''}>
		<span role="button" on:click={handleClick(tab.value)}>{tab.label}</span>
	</li>
    {/each}
</ul>

{#each tabs as tab}
	{#if activeTabValue == tab.value}
	<div class="box">
		<svelte:component this={tab.component} response_text={response_text} response_ir={response_ir} response_after_import={response_after_import} stdout_compiler={stdout_compiler} response_wasm={response_wasm}/>
	</div>
	{/if}
{/each}




<!--
<div class="output">
<h2>Output</h2>
{response_text}
</div>

<style>
    .output {
        height: 80px;
        margin-left: 10px;
    }
</style>

-->

<style>
    .box {
		margin-bottom: 10px;
		padding: 40px;
		border: 1px solid #dee2e6;
        border-radius: 0 0 .5rem .5rem;
        border-top: 0;
	}
  ul {
    display: flex;
    flex-wrap: wrap;
    padding-left: 0;
    margin-bottom: 0;
    list-style: none;
    border-bottom: 1px solid #dee2e6;
  }
	li {
		margin-bottom: -1px;
	}

  span {
    border: 1px solid transparent;
    border-top-left-radius: 0.25rem;
    border-top-right-radius: 0.25rem;
    display: block;
    padding: 0.5rem 1rem;
    cursor: pointer;
  }

  span:hover {
    border-color: #e9ecef #e9ecef #dee2e6;
  }

  li.active > span {
    color: #495057;
    background-color: #fff;
    border-color: #dee2e6 #dee2e6 #fff;
  }
</style>