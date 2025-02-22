/* eslint-disable @typescript-eslint/no-explicit-any */
import { extendTheme, type ThemeConfig } from "@chakra-ui/react"
import { mode } from "@chakra-ui/theme-tools"

const config: ThemeConfig = {
  initialColorMode: "dark",
  useSystemColorMode: true,
}

// extend the theme
const theme = extendTheme({
  config,

  styles: {
    global: (props: any) => ({
      body: {
        backgroundColor: mode("gray.500", "")(props),
      },
    }),
  },
})

export default theme