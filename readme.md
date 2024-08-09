# Introduction

This is the sample repo for my `fintech_devcon` 2024 talk titled `How to create a synthetic data generator from scratch using open source tools`. This is the completed repository but in the talk we go through building this step by step.

# Tools

1. Go
2. Bento
3. Postgres

# Setting up our local go environment

1. Create a new project and then create a folder to hold all of your go-code. Then open up a terminal window in that folder and download [Bento](https://warpstreamlabs.github.io/bento/docs/guides/getting_started/#install). Bento is a declarative streaming platform and will be the core of our pipeline.

2. If you don't already have Go, download go either by using brew, `brew install go`, or by going to the [Go website](https://go.dev/doc/install) and downloading Go from source. If you're using windows, you can use the windows installer with the `.msi` file or optionally download it from source.

3. Once you've installed Go, then you can verify the installation by running `go version`

# Setting up our local python environment

1. Create a new folder to hold your python code and then open up a terminal window in that folder. Create a virtual environment to hold your python dependencies by running : `python3 -m venv env`.

2. Then activate the environment by running: `source env/bin/activate`

3. Then download the following dependencies using `pip install`:

- pandas
- ctgan
- SQLAlchemy



# Deck link

[link](https://docs.google.com/presentation/d/1VljTYKMBZUNr7iu0j7Pu_Y16t4DoZ2ph/edit?usp=sharing&ouid=104644943787296427683&rtpof=true&sd=true)
