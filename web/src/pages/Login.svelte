<script lang="ts">
    const loginButtonSize =
        window.innerWidth -
        (3 * parseFloat(getComputedStyle(document.documentElement).fontSize));
</script>

<svelte:head>
    <script src="https://accounts.google.com/gsi/client" async defer></script>
    <script>
        function handleCredentialResponse(response) {
            document.cookie = `credential=${response.credential}`;

            window.location = "/";
        }
    </script>
</svelte:head>

<div id="root" class="container mt-6">
    <div class="is-flex-grow-1" />

    <section class="section is-large is-flex is-flex-direction-column mt-6">
        <h1 class="title">Tá na hora do show!</h1>
    </section>

    <section
        class="section is-flex-grow-2 is-flex is-flex-direction-column is-justify-content-center"
    >
        <div
            id="g_id_onload"
            data-client_id="698154937532-gj4c6ofq2t0q03ffv7vbimuf0cfrc1g8.apps.googleusercontent.com"
            data-auto_prompt="false"
            data-callback="handleCredentialResponse"
        >
            <div
                class="g_id_signin"
                data-type="standard"
                data-size="large"
                data-theme="outline"
                data-text="sign_in_with"
                data-shape="circle"
                data-width={loginButtonSize}
                data-logo_alignment="left"
            />
        </div>
    </section>

    <div class="is-flex-grow-1" />
</div>

<style>
    #root {
        display: flex;
        flex-direction: column;
        height: 100vh;
    }
</style>
