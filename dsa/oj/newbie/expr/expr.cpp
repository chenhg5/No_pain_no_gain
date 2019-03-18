#include <stdio.h>

using namespace std;

int main() {
    double a, b;

    char op;

    scanf("%lf %c %lf", &a, &op, &b);

    switch (op) {
        case '+':
            printf("%.0lf\n", a + b);
            break;
        case '-':
            printf("%.0lf\n", a - b);
            break;
        default:
            printf("%.0lf\n", a * b);
    }

    return 0;
}