<script setup>
import {onMounted, ref, watch} from "vue";
import {getCarouselDataListByCompany} from "@/api/carouselData";

const props = defineProps({
    companyName: String
})

const dataList1 = ref([])

watch(() => props.companyName, (newName) => {
    getCarouselDataListByCompany({company: newName}).then((res) => {
        dataList1.value = res.data.list
    })
})

function modifyImageLink(imageLink) {
    if (imageLink.indexOf("uploads/") === 0) {
        return `api/${imageLink}`
    } else {
        return imageLink
    }
}

function modifyRedirectLink(redirectLink) {
    if (redirectLink === "") {
        return "#"
    } else {
        return redirectLink
    }
}

function click(redirectLink) {
    if (redirectLink === "") {
    } else {
        window.location.replace(redirectLink)
    }
}
</script>

<template>
    <a-carousel autoplay>
        <template v-for="d in dataList1">
            <div>
                <a @click="click(d.redirectLink)"><img :src="modifyImageLink(d.imageLink)" alt=""></a>
            </div>
        </template>
    </a-carousel>
</template>

<style scoped>
img {
    width: 100%;
}
</style>