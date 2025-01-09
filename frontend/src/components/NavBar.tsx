import { Box, Container, Flex, Text, useColorMode, useColorModeValue } from '@chakra-ui/react'

function NavBar() {
  const { colorMode, toggleColorMode } = useColorMode()

  return (
    <Container maxW="900px">
      <Box
        bg={useColorModeValue("gray.400", "gray.700")}
        px={4}
        my={4}
        borderRadius="5"
      >
        <Flex
          justifyContent="center"
          alignItems="center"
          gap={3}
          display={{
            base: "none",
            sm: "flex"
          }}
        >
          <img src="/react.png" alt="logo" width={50} height={50} />
          <Text fontSize="40">+</Text>
          <img src="/go.png" alt="logo" width={40} height={40} />
          <Text fontSize="40">=</Text>
          <img src="/explode.png" alt="logo" width={50} height={50} />
        </Flex>
      </Box>
    </Container>
  )
}

export default NavBar