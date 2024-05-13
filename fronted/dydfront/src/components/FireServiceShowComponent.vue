<template>
    <v-container class="outerScreen">
        <v-card subtitle="以下是您的火化方案" title="您已经填写过属于您的火化方案" hover>
            <template v-slot:text>
                <div class="baseData">
                    <div class="rows">
                        <div>姓名：</div>
                        <div>年龄：</div>
                        <div>性别：</div>
                    </div>
                    <div class="rows">
                        <div>身份证号：</div>
                        <div>手机号：</div>
                        <div>火化日期：</div>
                    </div>
                    <div class="rows">
                        <div>住址：</div>
                    </div>
                </div>
            </template>
            <v-card-text>
                <div class="font-weight-bold ms-1 mb-2">已经选择的方案</div>

                <v-timeline align="start" density="compact">
                    <v-timeline-item dot-color="deep-purple-lighten-1" size="x-small">
                        <div class="selectText">殡仪馆：</div>
                    </v-timeline-item>
                    <v-timeline-item dot-color="deep-purple-lighten-1" size="x-small">
                        <div class="selectText">火化方式：</div>
                    </v-timeline-item>
                    <v-timeline-item dot-color="deep-purple-lighten-1" size="x-small">
                        <div class="selectText">墓地位置：</div>
                    </v-timeline-item>
                    <v-timeline-item dot-color="deep-purple-lighten-1" size="x-small">
                        <div class="selectText">骨灰盒样式：</div>
                    </v-timeline-item>
                </v-timeline>
            </v-card-text>
            <template v-slot:actions>
                <v-btn text="修改方案"></v-btn>
            </template>
        </v-card>
    </v-container>
</template>

<style scoped>
.outerScreen {
    flex: 50;

    .baseData {
        display: flex;
        flex-direction: column;
        justify-content: space-between;

        .rows {
            font-size: large;
            display: flex;
            flex-direction: row;
            justify-content: space-around;
            margin-top: 2rem;
        }
    }
}
</style>

<script setup>
import { ref } from 'vue';

let username = ref('')
</script>