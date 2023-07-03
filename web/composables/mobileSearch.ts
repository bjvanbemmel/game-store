export const useMobileSearch = () => useState<Boolean>('mobileSearch', () => false)
export const toggleMobileSearch = () => useMobileSearch().value = !useMobileSearch().value
