<script lang="ts">
    export let response_text : string;
    export let response_ir: string;
    export let response_after_import: string;
    export let stdout_compiler: string;
    export let response_wasm: string;
</script>

<div class="output-container">
<h2>Output</h2>
<span>Compiler Output</span>
<br>
<div class="compiler-output">
{stdout_compiler}
</div>

<span>
    Program Output
</span>
<br>

<div class="output">
{response_text}
</div>

</div>

<style>
    .output-container {
        height: 550px;
        width: 700px;
        margin-left: 10px;
    }

    .compiler-output {
        white-space: pre-line;
        overflow: auto;
        margin: 10px;
    }
    .output {
        margin-top: 10px;
    }
    span {
        color: #bbb;
	    white-space: nowrap;
    }
    span::before, span::after {
        flex: 1 1 auto;
        border-top: 1px solid #bbb;
        content: '';
    }
</style>