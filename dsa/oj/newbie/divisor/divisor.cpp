#include <cstdio>

int main() {

    int a, b, c = 1;

    scanf("%d%d", &a, &b);

// 暴力求解法  O(n)
//    if (a < b) {
//        for (int i = 1; i < a + 1; ++i) {
//            if (a % i == 0 && b % i == 0) {
//                c = i;
//            }
//        }
//    } else {
//        for (int i = 1; i < b + 1; ++i) {
//            if (a % i == 0 && b % i == 0) {
//                c = i;
//            }
//        }
//    }

    // 最大公约数求解办法：
    // 更相减损法
    // 辗转相除法 O(2logn) TODO: 证明

    // 辗转相除法
    while (c) {
        c = a % b;
        a = b;
        b = c;
    }

    // 设 x > y ,  且 x = r + y * c , 其中 r = x % y , c > 0 ;
    // x > r * (c+1) > 2 * r ;

    printf("%d\n", a);

    return 0;
}