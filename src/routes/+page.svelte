<script lang="ts">
    import type monaco from 'monaco-editor';
    import { onMount, onDestroy } from 'svelte';
    import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
    import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker';
    import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker';
    import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker';
    import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker';
    import OutputTabs from '$lib/output-tabs.svelte';
    import { config, grammar } from "$lib/monaco_def";

    type responseJson = {
    output: string,
    output_ir: string,
    output_after_imports: string,
    };

    let divEl: HTMLDivElement;
    let editor: monaco.editor.IStandaloneCodeEditor;
    let Monaco : typeof monaco;
    let response_text: string = "";
    let response_ir: string = "";
    let response_after_import: string = "";
    console.log(typeof(OutputTabs));
    onMount(async () => {
        self.MonacoEnvironment = {
            getWorker: function (_moduleId: any, label: string) {
                if (label === 'json') {
                    return new jsonWorker();
                }
                if (label === 'css' || label === 'scss' || label === 'less') {
                    return new cssWorker();
                }
                if (label === 'html' || label === 'handlebars' || label === 'razor') {
                    return new htmlWorker();
                }
                if (label === 'typescript' || label === 'javascript') {
                    return new tsWorker();
                }
                return new editorWorker();
                }
        };
        Monaco = (await import('monaco-editor'));
        //console.log("monaco ", Monaco)
        Monaco.languages.register({
            id: "cpoint",
        });

        Monaco.languages.onLanguage("cpoint", async () => {
            Monaco.languages.setLanguageConfiguration("cpoint", config);
            Monaco.languages.setMonarchTokensProvider("cpoint", grammar);
        });

        editor = Monaco.editor.create(divEl!, {
            value: ['import @std/print.cpoint', '', 'func main() {', '\tprintstr("Hello world!")', '}'].join('\n'),
            language: "cpoint",
            autoClosingBrackets: 'always',
        });

        /*return () => {
            editor.dispose();
        };*/
            
    });
    onDestroy(() => {
        Monaco?.editor.getModels().forEach((model) => model.dispose());
        editor?.dispose();
    });
    async function get_server_address(): Promise<string> {
        const localhost_string = "http://localhost:8080";
        const localhost_test = await fetch(`${localhost_string}/status`, {
            method: 'GET',
            headers: {
				'Content-Type': 'application/json'
			}
        });
        console.log(localhost_test)
        if (localhost_test.status == 200){
            return localhost_string;
        }

        // add then the real backend to test if it is up
        window.alert("Can't find server");
        return "";
    }
    async function run(){
        // communicate with API
        //console.log(editor)
        if (!editor){
            return null
        }
        console.log(editor.getValue())
        const server_address = await get_server_address();
        if (server_address == ""){
            return null;
        }
        const address = server_address + "/run-code"
        const body = JSON.stringify({
            code: editor.getValue(),
            should_return_ir: true,
            should_return_after_imports: true,
        });
        const response = await fetch(address, {
            method: 'POST',
            body: body,
            headers: {
				'Content-Type': 'application/json'
			}
        });
        const response_json: responseJson = await response.json();
        response_text = response_json.output;
        response_ir = response_json.output_ir;
        response_after_import = response_json.output_after_imports;
        console.log("response : ", response_text);
        //console.log("response ir : ", response_ir);
        return null;
    }
</script>


<h1>Welcome to the C. playground</h1>

<button on:click={() => run()}>Run</button>

<div class="container">

<div class="monaco-container" bind:this={divEl} />

<div class="output-tab">

<OutputTabs response_text={response_text} response_ir={response_ir} response_after_import={response_after_import}></OutputTabs>

</div>
<!--
<div class="output">
<h2>Output</h2>
{response_text}
</div>
-->


</div>

<style>
    .container {
        margin-top: 20px;
        display: flex;
    }
    .monaco-container {
        width: 55%;
        height: 750px;
    }
    .output, .output-tab {
        height: 750px;
        margin-left: 10px;
    }

    .output-tab {
        display: initial;
    }
    button {
        float: right;
        margin-right: 10px;
        color: #FFFFFF;
        background-color: black;
        border-color: white;
	    border-radius: 16px;
        padding-top: 6px;
        padding-bottom: 6px;
        padding-right: 8px;
        padding-left: 8px;
        box-shadow: none;
        text-decoration: none;
    }

</style>