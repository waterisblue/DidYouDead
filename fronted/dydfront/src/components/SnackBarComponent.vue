<template>
    <v-snackbar :timeout="2000" :color="props.snackColor" v-model="props.snackbar">
        {{ props.snackText }}

        <template v-slot:actions>
            <v-btn color="pink" variant="text" @click="snackbar = false">
                Close
            </v-btn>
        </template>
    </v-snackbar>
</template>

<script setup>
import { defineProps } from 'vue';

const props = defineProps({
    snackText: String,
    snackbar: Boolean,
    snackColor: String
})
</script>