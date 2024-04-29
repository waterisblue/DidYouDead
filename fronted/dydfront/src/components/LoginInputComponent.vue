<template>
    <div>
        <v-text-field
        hide-details="auto"
        :label="inputText"
        ></v-text-field>
    </div>
</template>

<style scoped lang="less">

</style>

<script setup>
    import { defineProps } from 'vue';
    let name = 'LoginInputComponent'

    const props = defineProps({
        inputText: String
    })
</script>