<script setup>
import { useUsersStore } from "@/stores/users";
import Avatar from "primevue/avatar";
import Badge from "primevue/badge";
import InputText from "primevue/inputtext";
import Message from "primevue/message";
import { computed, onBeforeUnmount, onMounted, ref } from "vue";
import { useRoute } from "vue-router";


// Mengambil route saat ini
const route = useRoute();

// Mengakses parameter 'id' dari URL
const roomId = computed(() => route.params.roomId);

const store = useUsersStore()
const username = store.userName;
const userId = store.userId;
const jwtToken = store.token;

const message = ref('');
const messages = ref([]);
let ws = null;

onMounted(() => {
    ws = new WebSocket(`ws://localhost:9000/ws/join-room/${roomId.value}?userId=${userId}&username=${username}&token=${jwtToken}`);
    ws.onopen = () => {
        console.log('WebSocket connection established');
    };

    ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        messages.value.push(data);
    };

    ws.onclose = () => {
        console.log('WebSocket connection closed');
    };

    ws.onerror = (error) => {
        console.error('WebSocket error:', error);
    }
});

function sendMessage() {
    if (message.value) {
        ws.send(message.value);
        message.value = '';
    }
}
onBeforeUnmount(() => {
    if (ws) {
        ws.close(); // Menutup WebSocket
        console.log('WebSocket connection closed on component unmount');
    }
});
</script>

<template>
    <div>
        <h1 class="mb-6">Room ID ------~ <Badge :value="roomId"></Badge></h1>
        <div v-for="(msg, index) in messages" :key="index" class="mb-3 flex flex-row items-center">
            <Message class="transition-colors" style="max-width: 25rem; min-width: 20rem;">{{ msg.content }}</Message>
            <Avatar icon="pi pi-user" class="ml-2" size="large" shape="circle" />
            <p>{{ msg.username }}</p>
        </div>
        <InputText v-model="message" :invalid="message === null" @keyup.enter="sendMessage"  placeholder="Type a message and press Enter"/>
    </div>

</template>