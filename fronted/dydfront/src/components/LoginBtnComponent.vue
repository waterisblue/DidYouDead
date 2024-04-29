<template>
    <div>
        <v-btn class="btnStyle" elevation="6" :color="btnColor">
            {{ btnText }}
        </v-btn>
    </div>
</template>

<style scoped>
    .btnStyle {
        width: 50vh;
    }
</style>

<script setup>
    import { defineProps } from 'vue';
    let name = 'LoginBtnComponent'

    const props = defineProps({
        btnText: String,
        btnColor: String
    })
</script>