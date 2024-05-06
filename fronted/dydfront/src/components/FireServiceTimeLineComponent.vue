<template>
    <v-timeline direction="horizontal">
      <v-timeline-item
      :dot-color="timeLineInfo == 'one' ? 'info' : ''"
      >
        <div>
          <div class="text-h6">殡仪馆介绍及选择</div>
        </div>
      </v-timeline-item>
  
      <v-timeline-item
      :dot-color="timeLineInfo == 'two' ? 'info' : ''"
      >
        <div>
          <div class="text-h6">填写信息</div>
        </div>
      </v-timeline-item>
  
      <v-timeline-item
      :dot-color="timeLineInfo == 'three' ? 'info' : ''"
      >
        <div>
          <div class="text-h6">火化方式</div>
        </div>
      </v-timeline-item>
      <v-timeline-item
      :dot-color="timeLineInfo == 'four' ? 'info' : ''"
      >
        <div>
          <div class="text-h6">骨灰盒样式</div>
        </div>
      </v-timeline-item>
      <v-timeline-item
      :dot-color="timeLineInfo == 'five' ? 'info' : ''"
      >
        <div>
          <div class="text-h6">墓地位置</div>
        </div>
      </v-timeline-item>
      <v-timeline-item
      :dot-color="timeLineInfo == 'six' ? 'info' : ''"
      >
        <div>
          <div class="text-h6">提交流程</div>
        </div>
      </v-timeline-item>
    </v-timeline>
  </template>

<script setup>
import { useTimeLineStore } from '@/stores/user';
import { storeToRefs } from 'pinia';

const timeLineStore = useTimeLineStore()

const {timeLineInfo} = storeToRefs(timeLineStore)

</script>