<template>
    <v-container class="designOuter">
        <div>
            <v-text-field hide-details="auto" label="资源名称" v-model="sourceName"></v-text-field>
        </div>
        <div>
            <v-text-field hide-details="auto" label="资源简介" v-model="supplySubtitle"></v-text-field>
        </div>
        <div>
            <v-textarea label="资源描述" v-model="sourceDetail"></v-textarea>
        </div>
        <div class="designDiv">
            <div>
                <v-file-input @change="pushFiles" label="上传资源图片" v-model="files"></v-file-input>
            </div>
            <div>
                <v-select label="资源类型" v-model="type" :items="['殡仪馆', '墓地', '火化方式', '骨灰盒']"></v-select>
            </div>
            <div>
                <v-text-field label="资源价格(元)" v-model="price"></v-text-field>
            </div>
        </div>
        <div class="sourceBtns">
            <v-btn class="btn" color="warning" @click="BackToList">
                返回
            </v-btn>
            <v-btn class="btn" color="info" @click="submitSource">
                提交
            </v-btn>
        </div>
        <SnackBarComponent :snackText="snackText" :snackbar="snackBar" snackColor="info" />
    </v-container>
</template>

<script setup>
import axios from '../axios';
import { ref } from 'vue'
import SnackBarComponent from './SnackBarComponent.vue';
import { useRouter } from 'vue-router';

let sourceName = ref('')
let sourceDetail = ref('')
let supplySubtitle = ref('')
let type = ref('')
let price = ref(0)
let files = ref(null)

// snackBar
let snackText = ref('')
let snackBar = ref(false)

const router = useRouter()

function submitSource() {
    let file = null
    if (files.value)
        file = files.value[0]
    const formData = new FormData()
    formData.append('imgfile', file)
    formData.append('sourcename', sourceName.value)
    formData.append('type', type.value)
    formData.append('subtitle', supplySubtitle.value)
    formData.append('price', price.value)
    formData.append('sourcedetail', sourceDetail.value)
    axios.post('/loginafter/addsupply', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    }).then(res => {
        snackBar.value = true
        snackText.value = "遗嘱保存成功！"
        BackToList()
    })
}

function BackToList(){
    router.push('/super/supply')
}
</script>

<style lang="less" scoped>
.designOuter {
    display: flex;
    width: 100%;
    height: 100%;
    flex-direction: column;
    justify-content: space-around;

    .sourceBtns {
        display: flex;
        flex-direction: row;
        justify-content: space-around;

        .btn {
            width: 25rem;
        }
    }
    .designDiv {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        div {
            width: 20rem;
        }
    }
}
</style>