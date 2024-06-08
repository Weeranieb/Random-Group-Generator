# Team Redistribution Project

This Go project redistributes members from a set of original teams into a new set of teams. The goal is to evenly distribute members while accommodating differences in the original and new team sizes. Random selection ensures a varied distribution each time the program runs.

## Features

- Evenly redistributes members from original teams to new teams.
- Ensures all new teams have a balanced number of members, with minor variations managed through random allocation.
- Supports dynamic resizing of teams based on member counts.

## Implementation Details

### Data Structures

1. **AllTeam**: A map representing the original teams and their members.
2. **NewTeam**: A map representing the new teams, initially empty.

### Redistribution Logic

1. **Member Counting**: The program counts current members in each new team to determine how many members are needed to balance the teams.
2. **Quota Calculation**:
   - Determines the number of members each new team should receive (quota).
   - Handles cases where the total number of members doesn't evenly divide among the new teams.
3. **Random Allocation**:
   - Randomly selects members from the original teams to allocate to new teams.
   - Ensures minor imbalances are corrected through random distribution of remaining members.

### Key Functions

- **Main Function**: Orchestrates the redistribution process, updates the `NewTeam` map, and prints the results.
- **Quota Management**: Balances the quota for each team based on the number of members required.
- **Random Selection**: Ensures a varied and fair distribution of members.

## Usage

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/your-username/team-redistribution.git
   ```
2. **Navigate to the Project Directory**:
   ```sh
   cd team-redistribution
   ```
3. **Run the Program**:
   ```sh
   go run main.go
   ```

## Example Output

The program prints the distribution of members in each new team:
-----> Team Alpha: 5 person(s)
Name1 A
Name3 B
Name4 C
...
-----> Team Bravo: 5 person(s)
Name2 A
Name5 B
Name6 C
...
-----> Team Charlie: 5 person(s)
Name3 A
Name1 B
Name2 C

## Future Enhancements

- **Configuration File**: Allow input teams and members to be defined in a configuration file.
- **Command-line Arguments**: Enable user-defined team names and sizes through command-line arguments.
- **Logging**: Add logging for better tracking of the redistribution process.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
