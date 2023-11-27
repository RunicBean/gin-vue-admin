<script setup lang="ts">

import Item from "./Item.vue";
defineProps<{
    dataList: Array<{
        imageLink: string,
        extLink: string,
        title: string
    }>,
    sliceNumbers: Array<number>
}>()

</script>

<template>
    <a-row>
        <a-col span="8">
            <Item
                :image-link="dataList[0].extLink"
                :image-url="dataList[0].imageLink"
                :title="dataList[0].title"
            />
        </a-col>
        <a-col><a-divider type="vertical" /></a-col>

        <a-col span="8">
            <template v-if="dataList.length > 1">
                <Item
                    :image-link="dataList[1].extLink"
                    :image-url="dataList[1].imageLink"
                    :title="dataList[1].title"
                />
            </template>

        </a-col>
        <a-col><a-divider type="vertical" /></a-col>
        <a-col span="8">
            <template v-if="dataList.length > 2">
                <Item
                    :image-link="dataList[2].extLink"
                    :image-url="dataList[2].imageLink"
                    :title="dataList[2].title"
                />
            </template>

        </a-col>
    </a-row>
</template>

<style scoped>
.ant-divider-vertical {
    height: 100%;
    margin: 0 -0.5px;
}
</style>