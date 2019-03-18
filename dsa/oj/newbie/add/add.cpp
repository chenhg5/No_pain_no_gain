#include <stdio.h>

using namespace std;

int main() {

    // int 的范围为：
    // long 的范围为：
    // long long 的范围为：
    // double 的范围为：
    // float 的范围为：
    long long a, b;

    // 建议用 strtoll, 检测输入准确性
    scanf("%lld%lld", &a, &b);
    printf("%lld", a + b);

    return 0;
}