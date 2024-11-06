import { ref, type Ref } from 'vue';
import { defineStore } from "pinia";
import { type GlobalThemeOverrides } from 'naive-ui';

interface GlobeTheme {
    asideBgColor: string;
    cardBgColor: string;
}

interface CustomThemeConfig {
    topLogoBgColor: string;
    topLogoTextColor: string;
    topHeaderBgColor: string;
    topHeaderTextColor: string;
    globeTheme: GlobeTheme;
    selfOverride: GlobalThemeOverrides;
}

interface Theme {
    bambooGreen: Ref<CustomThemeConfig>;
    contentPath: Ref<string>;
    menuSelected: Ref<string>;
    userPath: Ref<string>;
}

const useThemeStore = defineStore('themeStore', () => {
    // 通用部分
    const globeTheme = {
        asideBgColor: '#ffffff',
        cardBgColor: '#ffffff',
    };

    // 主题配置 若竹
    const bambooGreen = ref<CustomThemeConfig>({
        topLogoBgColor: '#5dac81',
        topLogoTextColor: '#bfe1e0',
        topHeaderBgColor: '#5dac81',
        topHeaderTextColor: '#ffffff',
        globeTheme,
        selfOverride: {
            common: {
                primaryColor: '#5dac81',
                primaryColorPressed: '#5dac81',
                bodyColor: '#eff2f7',
                modalColor: '#eff2f7',
                tableColor: '#eff2f7',
            },
            Button: {
                color: '#5dac81',
                colorFocusPrimary: '#4ca879',
                colorHoverPrimary: '#4ca879',
                colorPressedPrimary: '#4ca879',
                textColorPrimary: '#ffffff',
                textColorHoverPrimary: '#ffffff',
                textColorPressedPrimary: '#ffffff',
                textColorFocusPrimary: '#ffffff',
                colorTextPrimary: '#5dac81',
                colorTextHoverPrimary: '#4ca879',
                textColorPrimaryDisabled: '#ffffff',
            },
            Input: {
                borderHover: '1px solid #5dac81',
                borderFocus: '1px solid #5dac81',
            },
            Card: {
                colorEmbedded: '#ffffff',
            },
            Switch: {
                railColorActive: '#5dac81',
            },
            DataTable: {
                borderRadius: '3px',
            },
            Notification: {
                color: '#fff',
                bordered: '1px solid #5dac81',
            },
            Dropdown: {
                color: '#fff',
            },
        } as GlobalThemeOverrides,
    });

    // 当前访问的位置
    let userPath = ref<string>('/dashboard/summary');
    let contentPath = ref<string>('/admin/dashboard/summary');
    let menuSelected = ref<string>('dashboard');

    return {
        bambooGreen,
        contentPath,
        menuSelected,
        userPath,
    };
}, {
    persist: true,
});

export default useThemeStore;