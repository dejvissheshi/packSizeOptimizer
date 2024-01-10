import sys
import pulp

def optimize_packet_distribution(xi, W):
    # Number of packet types
    n = len(xi)

    # First Optimization: Minimize sum of wi * xi - W
    problem1 = pulp.LpProblem("Minimize_Packet_Size", pulp.LpMinimize)

    # Variables for wi (quantity of each packet type)
    wi = [pulp.LpVariable(f'w{i}', lowBound=0, cat='Integer') for i in range(n)]

    # Objective function for the first optimization
    problem1 += pulp.lpSum([wi[i] * xi[i] for i in range(n)]) - W

    # Constraint: sum of wi * xi >= W
    problem1 += pulp.lpSum([wi[i] * xi[i] for i in range(n)]) >= W

    # Solve the first optimization problem
    problem1.solve(pulp.PULP_CBC_CMD(msg=False))

    if pulp.LpStatus[problem1.status] != 'Optimal':
        return "No feasible solution found for the first optimization"

    # Minimum value of sum of wi * xi - W
    min_value = pulp.value(problem1.objective)

    # Second Optimization: Minimize sum of wi among solutions with the same minimum value
    problem2 = pulp.LpProblem("Minimize_Number_of_Packages", pulp.LpMinimize)

    # Objective function for the second optimization
    problem2 += pulp.lpSum([wi[i] for i in range(n)])

    # Constraint: sum of wi * xi - W should be equal to the minimum value found
    problem2 += pulp.lpSum([wi[i] * xi[i] for i in range(n)]) - W == min_value

    # Solve the second optimization problem
    problem2.solve(pulp.PULP_CBC_CMD(msg=False))

    if pulp.LpStatus[problem2.status] != 'Optimal':
        return "No feasible solution found for the second optimization"

    # Return the optimized quantities wi for each packet size
    return [int(wi[i].varValue) for i in range(n)]

def main():
    try:
        numbers = [int(arg) for arg in sys.argv[1:]]
    except ValueError:
        sys.exit(1)

    if not numbers:
       sys.exit(1)

    # Extract the new number from the command line arguments
    try:
        new_number = int(numbers.pop())
    except (ValueError, IndexError):
        sys.exit(1)

    xi = numbers
    W = new_number
    return optimize_packet_distribution(xi, W)

if __name__ == "__main__":
    result = main()
    print(result)