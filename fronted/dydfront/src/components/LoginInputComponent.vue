<template>
    <div>
        <v-text-field
        hide-details="auto"
        :label="inputText"
        @input="$emit('sendInputData', inputValue)"
        v-model="inputValue"
        :rules="thisRules"
        ></v-text-field>
    </div>
</template>

<style lang="less" scoped >

</style>

<script setup>
    import { ref } from 'vue';
    import { defineProps } from 'vue';

    let name = 'LoginInputComponent'
    let inputValue = ref('')
    const props = defineProps({
        inputText: String,
        thisRules: Object
    })
</script>