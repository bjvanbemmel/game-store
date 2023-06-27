<template>
    <div>
        <Line
            id="playercount"
            :options="chartOptions"
            :data="chartData"
            style="height: 12rem; width: 100%; position: relative;"
        />
    </div>
</template>

<script setup lang="ts">
import { Line } from 'vue-chartjs'
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Filler,
    Legend
} from 'chart.js'
import Game from '~/types/game'
import { defineProps } from 'vue'

ChartJS.register(CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Filler,
    Legend
)

const props = defineProps<{
    game: Game
}>()

const chartData = {
    labels: Array<number>(),
    datasets: [{
        label: 'Players',
        backgroundColor: "rgba(71, 183,132,.5)",
        borderColor: "#47b784",
        borderWidth: 3,
        fill: true,
        tension: 0.1,
        data: Array<number>(),
    }],
}

const chartOptions = {
    responsive: true,
}

chartData.labels = props.game.player_count.map(x => x.hour)
chartData.datasets[0].data = props.game.player_count.map(x => x.count)
</script>
