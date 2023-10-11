<script lang="ts">
    export let response_text: string;
    export let response_ir: string;
    export let response_after_import: string;
    export let stdout_compiler: string;
    export let response_wasm: string;
</script>

<div class="output">
<h2>IR</h2>
{response_ir}
</div>

<style>
    .output {
        height: 550px;
        width: 700px;
        margin-left: 10px;
        white-space: pre-line;
        overflow: auto;
    }
</style>