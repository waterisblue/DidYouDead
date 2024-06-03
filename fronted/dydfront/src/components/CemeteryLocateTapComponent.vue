<template>
  <div class="outerScreen">
    <v-card class="card" max-width="240" min-width="300">
      <img height="200px"   :src=tapImage cover class="tabimages" />
      <v-card-title>
        {{ props.tapName }}
      </v-card-title>
      <v-card-subtitle>
        {{ props.tapSubTitle }}
      </v-card-subtitle>
      <v-card-actions>
        <v-btn color="orange-lighten-2" text="选择" @click="checkUrl('/fireservice/submit')"></v-btn>
        <v-spacer></v-spacer>
        <!-- <v-btn :icon="show ? 'mdi-chevron-up' : 'mdi-chevron-down'" @click="show = !show"></v-btn>  -->
      </v-card-actions>
      <v-expand-transition>
        <div v-show="show">
          <v-divider></v-divider>
          <v-card-text>
            {{ props.tapDetail }}
          </v-card-text>
        </div>
      </v-expand-transition>
    </v-card>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { useFireServiceChoose } from '@/stores/user';
const { getService, setService } = useFireServiceChoose()
let service = getService()

let show = ref(false)

const router = useRouter()
function checkUrl(path) {
  service.Cemetery = props.tapName
  service.CemeteryId = props.tapId
  setService(service)
  router.push(path)
}

const props = defineProps({
  tapImage: String,
  tapName: String,
  tapDetail: String,
  tapSubTitle: String,
  tapId: Int16Array
})
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

    .tabimages {
      display: block;
      margin: auto;
    }
  }
}

.outerScreen:hover {
  /* flex: 1.5;  */
  opacity: 0.9; 
  transition: 1s; 
  transform: scale(1.1); 
}
/* .modeCard{
  transition: all 0.5s;
}

.modeCard:hover{ //鼠标悬停时激活
  opacity: 0.2; //透明度
  transition: 1s; //动画过渡的更加顺滑
  transform: scale(1.1); //放大倍数
} */

</style>