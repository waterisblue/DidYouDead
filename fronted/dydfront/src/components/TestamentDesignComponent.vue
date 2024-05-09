<template>
    <v-container class="designOuter">
        <div>
            <v-text-field hide-details="auto" label="遗嘱名" v-model="testamentName"></v-text-field>
        </div>
        <div>
            <v-textarea label="遗嘱内容" v-model="testamentDetail"></v-textarea>
        </div>
        <div>
            <v-file-input @change="pushFiles" label="上传遗嘱证明文件" v-model="files"></v-file-input>
        </div>
        <div>
            <v-select label="遗嘱形式" v-model="testamentStyle" :items="['自书遗嘱', '代书遗嘱', '打印遗嘱', '公证遗嘱']"></v-select>
        </div>
        <div class="testamentBtns">
            <v-btn class="btn" color="error">
                清空
            </v-btn>
            <v-btn class="btn" color="info" @click="submitTestament">
                提交
            </v-btn>
        </div>
    </v-container>
</template>

<script setup>
import axios from '../axios';
import { ref } from 'vue'

let testamentDetail = ref('')
let testamentStyle = ref('')
let testamentName = ref('')
let files = ref(null)

// function pushFiles(file){
//     if(!file) return

//     files.value = file
// }
function submitTestament() {
    console.log(testamentDetail.value)
    console.log(testamentStyle.value)
    let file = files.value[0]
    const formData = new FormData()
    formData.append('testamentJustifyFile', file)
    formData.append('testamentDetail', testamentDetail.value)
    formData.append('testamentName', testamentName.value)
    formData.append('testamentStyle', testamentStyle.value)
    axios.post('/loginafter/uploadtestment', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    }).then(res => {
        console.log(res)
    })
}
</script>

<style lang="less" scoped>
.designOuter {
    display: flex;
    width: 100%;
    height: 100%;
    flex-direction: column;
    justify-content: space-around;

    .testamentBtns {
        display: flex;
        flex-direction: row;
        justify-content: space-around;

        .btn {
            width: 25rem;
        }
    }
}
</style>