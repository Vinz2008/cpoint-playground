<script lang="ts">
	export let menuOpen: boolean;
</script>

<button on:click class="dropbtn">{!menuOpen ? "Show Menu" : "Close Menu"}</button>

<style>
    button {
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