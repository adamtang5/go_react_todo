import { Box, Container, Flex, useColorMode, useColorModeValue } from '@chakra-ui/react'

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
          
        </Flex>
      </Box>
    </Container>
  )
}

export default NavBar