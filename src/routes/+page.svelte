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
    import { Pane, Splitpanes } from 'svelte-splitpanes';
    //import Dropdown from '$lib/dropdown.svelte';
    import DropdownButton from '$lib/dropdown-button.svelte';
    //import { checkbox_Output, checkbox_IR, checkbox_After_imports } from "$lib/dropdown-store"
    type responseJson = {
        stdout_compiler: string
        output: string,
        output_ir: string,
        output_wasm: string,
        output_after_imports: string,
    };
    let menuOpen: boolean = false;

    let checkbox_Output: boolean = true;
    let checkbox_IR: boolean = false;
    let checkbox_After_imports: boolean = false;
    let checkbox_Wasm: boolean = false;


    let divEl: HTMLDivElement;
    let editor: monaco.editor.IStandaloneCodeEditor;
    let Monaco : typeof monaco;
    let response_text: string = "";
    let response_ir: string = "";
    let response_after_import: string = "";
    let response_wasm: string = "";
    let stdout_compiler: string = "";
    console.log(typeof(OutputTabs));

    const STORAGE_KEY = 'theme';
    const DARK_PREFERENCE = '(prefers-color-scheme: dark)';
    const THEMES = {
        DARK: 'dark',
        LIGHT: 'light',
    };

    const toggleTheme = () => {
    const stored = localStorage.getItem(STORAGE_KEY);

    if (stored) {
      // clear storage
      localStorage.removeItem(STORAGE_KEY);
    } else {
      // store opposite of preference
      localStorage.setItem(STORAGE_KEY, prefersDarkThemes() ? THEMES.LIGHT : THEMES.DARK);
    }

    // TODO: apply new theme
    };

    const prefersDarkThemes = () => window.matchMedia(DARK_PREFERENCE).matches;



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
            should_return_ir: checkbox_IR,
            should_return_after_imports: checkbox_After_imports,
            should_return_wasm: checkbox_Wasm,
        });
        const response = await fetch(address, {
            method: 'POST',
            body: body,
            headers: {
				'Content-Type': 'application/json'
			}
        });
        const response_json: responseJson = await response.json();

        if (checkbox_Output){
            response_text = response_json.output;
        } else {
            response_text = "";
        }

        if (checkbox_IR){
            response_ir = response_json.output_ir;
        } else {
            response_ir = "";
        }

        if (checkbox_After_imports){
            response_after_import = response_json.output_after_imports;
        } else {
            response_after_import = "";
        }
        if (checkbox_Wasm){
            response_wasm = response_json.output_wasm;
        } else {
            response_wasm = "";
        }
        stdout_compiler = response_json.stdout_compiler;
        console.log("response : ", response_text);
        //console.log("response ir : ", response_ir);
        //console.log(stdout_compiler);
        return null;
    }
</script>


<h1>Welcome to the C. playground</h1>

<DropdownButton on:click={() => menuOpen = !menuOpen} {menuOpen}/>
    
{#if menuOpen}
<div id="dropdown"  class="dropdown-content">
    <label>
        <input type="checkbox" checked={checkbox_Output} on:change={() => {checkbox_Output = !checkbox_Output}}/>
            Output
    </label>
    
    <br>
    
    <label>
        <input type="checkbox" checked={checkbox_IR} on:change={() => {checkbox_IR = !checkbox_IR}}/>
            LLVM IR
    </label>
    
    <br>
    
    <label>
        <input type="checkbox" checked={checkbox_After_imports} on:change={() => {checkbox_After_imports = !checkbox_After_imports}}/>
            After Imports
    </label>

    <br>

    <label>
        <input type="checkbox" checked={checkbox_Wasm} on:change={() => {checkbox_Wasm = !checkbox_Wasm}}/>
            WASM
    </label>
</div>
{/if}


<button on:click={() => run()}>Run</button>

<div class="container">

<Splitpanes>

<Pane minSize={15}>
    
<div class="monaco-container" bind:this={divEl} />

</Pane>

<Pane>

<div class="output-tab">

<OutputTabs response_text={response_text} response_ir={response_ir} response_after_import={response_after_import} response_wasm={response_wasm} stdout_compiler={stdout_compiler}></OutputTabs>

</div>

</Pane>

</Splitpanes>
<!--
<div class="output">
<h2>Output</h2>
{response_text}
</div>
-->


</div>

<button on:click={toggleTheme}>Dark Mode</button>


<svelte:head>
	<link rel="preconnect" href="https://fonts.bunny.net">
    <link href="https://fonts.bunny.net/css?family=source-code-pro:400" rel="stylesheet" />
</svelte:head>

<style>
    :root {
        font-family: 'Source Code Pro', monospace;
    }
    .container {
        margin-top: 20px;
        display: flex;
    }
    .monaco-container {
        /*width: 55%;*/
        width: 100%;
        height: 750px;
    }
    .output, .output-tab {
        height: 750px;
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

    #dropdown {
        position: absolute;
        z-index: 1;
        background-color: white;
        border: 1px solid black;
    }

</style>