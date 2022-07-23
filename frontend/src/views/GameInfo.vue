<template>
	<v-container>
		<h2>{{ gameTitle }} {{ $route.params.id }}</h2>

		<v-card>
			
			<v-img width="200" :src="image"></v-img>
			<router-link :to="{ name: 'info', params: { id: '2' }}">To another game</router-link>
	<!-- Title -->
	<!-- Ratings -->
	<v-divider> </v-divider>
		<v-card-actions class="pa-3"> 
			 Rate {{ gameTitle }}
			<v-spacer></v-spacer>
			<span class="pr-2">
				{{ rating }}
			</span>
			<v-rating
				v-model="rating"
				background-color="#8A8D8F"
				color="#00263A"
				half-increments
				dense
				hover
				size="18"
			></v-rating> 
		</v-card-actions>
	<!-- Overview: Image -->
	<!-- Download game -->
		</v-card>
	</v-container>
</template>

<script>
export default {
	name: 'GameInfo',

	data: () => ({
		gameTitle: 'Big Blue Game',
		rating: 3,
		image: 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg'
	}),

	computed: {

	},

	methods: {

	}
};

</script>