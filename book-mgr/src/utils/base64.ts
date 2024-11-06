// 明文转 Base64
const encodeToBase64 = (plainText: string): string => {
    // 将明文转换为 Base64 编码
    return btoa(plainText);
}

export  {
    encodeToBase64,
}