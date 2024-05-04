<template>
    <v-container class="designOuter">
        <div>
            <v-textarea label="遗嘱内容"></v-textarea>
        </div>
        <div>
            <v-file-input label="上传遗嘱证明文件"></v-file-input>
        </div>
        <div>
            <v-select label="遗嘱形式"
                :items="['自书遗嘱', '代书遗嘱', '打印遗嘱', '公证遗嘱']"></v-select>
        </div>
        <div class="testamentBtns">
            <v-btn class="btn" color="error">
                清空
            </v-btn>
            <v-btn class="btn" color="info">
                提交
            </v-btn>
        </div>
    </v-container>
</template>

<script>

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