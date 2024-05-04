<template>
  <div class="outerScreen">
  <v-card class="card" max-width="344" min-width="344">
    <v-img height="200px" src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg" cover></v-img>
    <v-card-title>
      Top western road trips
    </v-card-title>
    <v-card-subtitle>
      1,000 miles of wonder
    </v-card-subtitle>
    <v-card-actions>
      <v-btn color="orange-lighten-2" text="选择"></v-btn>
      <v-spacer></v-spacer>
      <v-btn :icon="show ? 'mdi-chevron-up' : 'mdi-chevron-down'" @click="show = !show"></v-btn>
    </v-card-actions>
    <v-expand-transition>
      <div v-show="show">
        <v-divider></v-divider>
        <v-card-text>
          I'm a thing. But, like most politicians, he promised more than he could deliver. You won't have time for
          sleeping, soldier, not with all the bed making you'll be doing. Then we'll go with that data file! Hey, you
          add a one and two zeros to that or we walk! You're going to do his laundry? I've got to find a way to escape.
        </v-card-text>
      </div>
    </v-expand-transition>
  </v-card>
</div>
</template>

<script setup>
import { ref } from 'vue';

let show = ref(true)
</script>

<style scoped>
.outerScreen {
  flex: 1;
  transition: all 1s;
  display: flex;
  align-items: center;
  justify-content: center;
  .card {
    transition: all 1s;
  }
}

.outerScreen:hover {
  flex: 1.5;
}
</style>