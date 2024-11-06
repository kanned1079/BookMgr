export let formatDate = (dateString: string) => {
    const date = new Date(dateString);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

export function formatTimestamp(timestamp: number): string {
    // 如果时间戳是秒级的，乘以1000转换为毫秒级
    if (timestamp.toString().length === 10) {
        timestamp *= 1000;
    }

    const date = new Date(timestamp);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以要加1
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

