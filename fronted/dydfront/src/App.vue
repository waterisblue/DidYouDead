<template>
  <v-layout class="rounded rounded-md">
    <v-navigation-drawer color="grey-darken-3" permanent>
      <v-list
          density="compact"
          nav
        >
          <v-list-item prepend-icon="mdi-view-dashboard" title="Home" value="home"></v-list-item>
          <v-list-item prepend-icon="mdi-forum" title="About" value="about"></v-list-item>
        </v-list>
    </v-navigation-drawer>

    <v-app-bar
      :order="order"
      color="grey-darken-4"
      title="Did You Dead?"
      flat
    >
    </v-app-bar>

    <v-main class="d-flex align-center justify-center" style="min-height: 300px;">
      Main Content
    </v-main>
  </v-layout>
</template>
<script setup>
  import {ref} from 'vue'

  const order = ref(-1)

</script>