<script lang="ts">
    import type monaco from 'monaco-editor';
    import { onMount, onDestroy } from 'svelte';
    import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
    import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker';
    import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker';
    import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker';
    import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker';

    import { config, grammar } from "$lib/monaco_def";

    let divEl: HTMLDivElement;
    let editor: monaco.editor.IStandaloneCodeEditor;
    let Monaco : typeof monaco;
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
            value: ['import @std/print.cpoint', '', 'func test() {', '\tprintstr("Hello world!")', '}'].join('\n'),
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

    function run(){
        // communicate with API
        // write it in Go ? in Rust ? In Js Runtime (Deno/Bun/Nodejs) ? 
        return null;
    }
</script>


<h1>Welcome to the C. playground</h1>


<button on:click={run()}>Run</button>
<div class="container" bind:this={divEl} />

<style>
    .container {
        width: 50%;
        height: 1000vh;
    }
</style>