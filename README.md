# JobTrack

JobTrack is a lightweight CLI tool that helps you track your job applications, built with Go and SQLite.
It is my first Go project (yay!).

## 🚀 Installation

### 1️⃣ Build from Source (Recommended if you have `Go` and `make`)

#### **Requirements:**

- **Go** installed (`go version` to check)
- **Make** installed (`make -v` to check)

#### **Install via Make** (Recommended)

```sh
make install
```

✅ This will:

- Build JobTrack
- Install it to /usr/local/bin
- Install man pages (man jobtrack, man jobtrack-create etc)

After make install, you will only need to move the completion script for your shell manually as shown [below](#completions).

#### **Manual Build Without Installation**

```sh
make build
```

✅ This will only build JobTrack (to an executable named jobtrack)

You will need to manually move the binary, completion script and man pages as shown [below](#manual).

### 2️⃣ Download Prebuilt Binaries (No Go Required)

If you do not want to build from source, download a prebuilt release.

#### Steps: <span id="manual"></span>

1. Go to the [Releases](https://github.com/valentino7504/jobtrack/releases) page.
2. Download the .zip file for your OS and architecture.
3. Extract and move the executable to a directory in your PATH:

```sh
sudo mv jobtrack /usr/local/bin/
```

4. Move man pages <span id="man-pages"></span>

```sh
sudo mv man/* /usr/share/man/man1/
mandb  # Update the man page database
```

5. Move autocompletions (if needed):<span id="completions"></span>

- Bash

```sh
sudo mv completions/jobtrack.bash /etc/bash_completion.d/
```

- Zsh

```sh
sudo mkdir /usr/local/share/zsh/site-functions/_jobtrack
sudo mv completions/jobtrack.zsh /usr/local/share/zsh/site-functions/_jobtrack
```

- Fish

```sh
sudo mv completions/jobtrack.fish /etc/fish/completions/
```

After moving the completion scripts, please restart your shell for changes to be effected.

## 📌 Usage

### Basic Commands

#### 1️⃣ Create a new job entry

Adds a new job to the database.

```sh
jobtrack create --company="Google" --position="Software Engineer"
```

###### Options:

- `--company` (required): Name of the company.
- `--position` (required): Job title you're applying for.
- `--status`: Application status (e.g., Applied, Interview, Offer).
- `--applied`: Date applied (YYYY-MM-DD).
- `--location`: Job location.
- `--salary-range`: Salary expectation.
- `--job-posting-url`: Link to the job posting.

#### 2️⃣ List jobs

Displays job applications stored in the database.

```sh
jobtrack list
```

###### Filtering Options:

- `--id`: Show a specific job by ID.
- `--status`: Show jobs with a specific status (e.g., Applied, Interview, Offer).
- `--after`: Show jobs applied to **after** a date (YYYY-MM-DD).
- `--before`: Show jobs applied to **before** a date (YYYY-MM-DD).

#### 3️⃣ Update a job entry

Modify an existing job entry.

```sh
jobtrack update --id=3 --status="Offer"
```

###### Filtering Options:

- `--id` (required): The ID of the job to update.
- Other flags (`--status`. `--company`, `--position`) to update their respective fields.

#### 3️⃣ Update a job entry

Modify an existing job entry.

```sh
jobtrack update --id=3 --status="Offer"
```

###### Filtering Options:

- `--id` (required): The ID of the job to update.
- Other flags (`--status`. `--company`, `--position`) to update their respective fields.

#### 4️⃣ Delete a job entry

Delete a job from the database using its ID.

```sh
jobtrack delete --id=5
```

⚠️ **This action is permanent!**

#### 5️⃣ Export jobs to JSON or CSV

Saves job applications to a file.

```sh
jobtrack export --format=json --output=jobs.json
```

###### Options:

- `--format` or `-f`: Choose `json` (default) or `csv`.
- `--output` or `-o`: Specify output file (prints to stdout by default).

**CSV export example**

```sh
jobtrack export --format=csv --output=jobs.csv
```

##### Exported JSON example (jobs.json)

```json
[
  {
    "company": "Reddit",
    "position": "SDE",
    "status": "Interview",
    "location": "Remote",
    "salary_range": "400k",
    "job_posting_url": "https://reddit.com/jobs",
    "applied_at": "2025-03-22T00:00:00Z",
    "created_at": "2025-03-22T12:59:27Z",
    "updated_at": "2025-03-22T12:59:27Z",
    "id": 1
  },
  {
    "company": "Amazon",
    "position": "SDE Intern",
    "status": "Applied",
    "location": "Remote",
    "salary_range": "200k",
    "job_posting_url": "https://amazon.com/jobs",
    "applied_at": "2025-03-22T00:00:00Z",
    "created_at": "2025-03-22T12:59:34Z",
    "updated_at": "2025-03-22T12:59:34Z",
    "id": 2
  }
]
```

##### Exported CSV example (jobs.csv)

```csv
Company,Position,Status,Location,SalaryRange,JobPostingURL,AppliedAt,CreatedAt,UpdatedAt
Reddit,SDE,Interview,Remote,400k,https://reddit.com/jobs,2025-03-22,2025-03-22 12:59:27,2025-03-22 12:59:27
Amazon,SDE Intern,Applied,Remote,200k,https://amazon.com/jobs,2025-03-22,2025-03-22 12:59:34,2025-03-22 12:59:34
```

#### 6️⃣ Import Jobs from JSON or CS

Loads job applications from a properly formatted file to the databaase.

From JSON:

```sh
jobtrack import jobs.json
```

From CSV:

```sh
jobtrack import jobs.csv
```

Please ensure the file is formatted correctly, I have **not** implemented checks for that and your installation might break.

## 📄 Man Pages

After installation, you can view the man pages using:

```sh
man jobtrack
man jobtrack-create
man jobtrack-list
man jobtrack-update
man jobtrack-delete
man jobtrack-import
man jobtrack-export
```

## 🗑️ Uninstallation

- If you wish to uninstall via `make`:

```sh
make uninstall
```

- If you wish to uninstall manually:

```sh
sudo rm -f /usr/local/bin/jobtrack
sudo rm -f /usr/share/man/man1/jobtrack*
rm -rf $HOME/.local/share/jobtrack
```

You can then remove the completion scripts from the [directories](#completions) where you installed them.

## 🤝 Contributions

1. Fork the repo
2. Clone your fork
3. Create a new branch
4. Make changes and commit
5. Push changes and make a PR

I'd be happy to review and merge any PR's that add some useful functionality.

## 🧾 License

JobTrack is released under the MIT license. See [LICENSE](./LICENSE).
