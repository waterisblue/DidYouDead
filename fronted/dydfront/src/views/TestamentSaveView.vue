<template>
    <v-container fluid class="fullScreen">
        <div class="toolBar" @click="checkUrl('/bubble')">
            back
        </div>
        <div class="testamentInput">
            <div class="testamentText">
                <v-textarea label="遗嘱" variant="solo-filled" style="min-height: 80vh;"></v-textarea>
            </div>
            <div class="testamentFiles">

            </div>
        </div>

    </v-container>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter()

function checkUrl(path){
    router.push(path)
}
</script>

<style scoped>
.fullScreen {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;

    .toolBar {
        flex: 1;
    }

    .testamentInput {
        flex: 30;
        display: flex;
        flex-direction: row;
        justify-content: center;

        .testamentText {
            flex: 1;
            background-color: aliceblue
        }

        .testamentFiles {
            flex: 1;
            background-color: antiquewhite;
        }
    }

}
</style>