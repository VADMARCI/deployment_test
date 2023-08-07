import argparse, os

def replace_using_env_vars(template, output):
    with open(template) as template:
        expanded = os.path.expandvars(template.read())
        with open(output, 'w') as out:
            out.write(expanded)

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('template')
    parser.add_argument('output')
    args = parser.parse_args()

    replace_using_env_vars(args.template, args.output)
