#include <iostream>

int main()
{
    const float pi = 3.124;
    int radius = 14;
    float area;

    area = pi * radius * radius;

    std::cout << "The Area is: " << area << "\n";

    return 0;
}