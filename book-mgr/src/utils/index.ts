import { h } from 'vue'
import type { Component } from 'vue'
import { NIcon } from 'naive-ui'

function renderIcon(icon: Component) {
    return () => {
        return h(NIcon, null, {
            default: () => h(icon)
        })
    }
}

export default renderIcon