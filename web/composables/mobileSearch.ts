export const useMobileSearch = () => useState<boolean>('mobileSearch', () => false)
export const toggleMobileSearch = () => useMobileSearch().value = !useMobileSearch().value
