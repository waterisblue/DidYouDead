import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
  }

  return { count, doubleCount, increment }
})

export const useTimeLineStore = defineStore('timeline', () => {
  const timeLineInfo = ref("one");

  function setTimeLineInfo(path){
      timeLineInfo.value = path
    }


  return {timeLineInfo, setTimeLineInfo}
}
)

export const useUserStore = defineStore('user', () => {
  const token = ref("")

  function setToken(token){
    this.token = token
  }

  return { token, setToken }
})