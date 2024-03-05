<template>
  <v-layout class="rounded rounded-md">
    <v-navigation-drawer color="grey-darken-3" permanent>
      <v-list
          density="compact"
          nav
        >
          <v-list-item prepend-icon="mdi-view-dashboard" title="Index" value="index" @click="checkUrl('/')"></v-list-item>
          <v-list-item prepend-icon="mdi-forum" title="Chart" value="chart" @click="checkUrl('/chart')"></v-list-item>
          <v-list-item prepend-icon="mdi-forum" title="Table" value="table" @click="checkUrl('/table')"></v-list-item>
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
      <RouterView />
    </v-main>
  </v-layout>
</template>
<script setup>
  import {ref} from 'vue'
  import { RouterView } from 'vue-router';
  import router from './router';

  const order = ref(-1)
  
  // 用于vue-router
  function checkUrl(path){
    router.push(path)
  }

</script>