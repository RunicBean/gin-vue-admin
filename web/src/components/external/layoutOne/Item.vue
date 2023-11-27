<script setup lang="ts">
defineProps([
    "imageLink",
    "imageUrl",
    "title"
])
</script>

<template>
    <a :href="imageLink">
        <a-flex vertical class="item-container" align="center" justify="space-between" gap="small">
            <img :src="imageUrl" alt="">
            <h1>{{title}}</h1>
        </a-flex>
    </a>
</template>

<style scoped>
.item-container {
    padding: 15px 8px;
}
img {
    width: 40%;
}

h1 {
    color: #1a1a1a;
    font-size: 14px;
}
</style>