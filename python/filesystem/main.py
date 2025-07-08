from pathlib import Path
from typing import List, Union, Optional
import os
from datetime import datetime


class Tree:
    def __init__(self, root: 'Node'):
        self.root = root

    def __str__(self):
        return str(self.root)

    def find_files_by_extension(self, extension: str) -> List['File']:
        """Find all files with given extension in the tree"""
        files = []

        def _search(node: 'Node'):
            if isinstance(node, File) and node.path.suffix.lower() == extension.lower():
                files.append(node)
            elif isinstance(node, Dir):
                for child in node.children:
                    _search(child)

        _search(self.root)
        return files

    def find_largest_files(self, count: int = 10) -> List['File']:
        """Find the largest files in the tree"""
        files = []

        def _collect_files(node: 'Node'):
            if isinstance(node, File):
                files.append(node)
            elif isinstance(node, Dir):
                for child in node.children:
                    _collect_files(child)

        _collect_files(self.root)
        return sorted(files, key=lambda f: f.size, reverse=True)[:count]


class Node:
    def __init__(self, path: Path, name: str = None):
        self.path = path
        self.name = name or path.name

    def __str__(self):
        return f"{self.__class__.__name__}: {self.name}"

    @property
    def last_modified(self) -> datetime:
        """Get last modification time"""
        return datetime.fromtimestamp(self.path.stat().st_mtime)

    @property
    def permissions(self) -> str:
        """Get file permissions as string"""
        stat = self.path.stat()
        return oct(stat.st_mode)[-3:]


class Dir(Node):
    def __init__(self, path: Path, name: str = None, children: List['Node'] = None):
        super().__init__(path, name)
        self.children = children or []

    @property
    def total_size(self) -> int:
        """Calculate total size of directory including all children"""
        return sum(
            child.size if isinstance(child, File) else child.total_size
            for child in self.children
        )

    @property
    def file_count(self) -> int:
        """Count total number of files in directory and subdirectories"""
        count = 0
        for child in self.children:
            if isinstance(child, File):
                count += 1
            elif isinstance(child, Dir):
                count += child.file_count
        return count

    @property
    def dir_count(self) -> int:
        """Count total number of subdirectories"""
        count = 0
        for child in self.children:
            if isinstance(child, Dir):
                count += 1 + child.dir_count
        return count

    def add_child(self, child: 'Node'):
        """Add a child node to this directory"""
        self.children.append(child)

    def get_child(self, name: str) -> Optional['Node']:
        """Get child by name"""
        for child in self.children:
            if child.name == name:
                return child
        return None

    def __str__(self):
        return f"Dir: {self.name} ({len(self.children)} items, {self._format_size(self.total_size)})"

    def _format_size(self, size: int) -> str:
        """Format size in human-readable format"""
        for unit in ['B', 'KB', 'MB', 'GB', 'TB']:
            if size < 1024.0:
                return f"{size:.1f}{unit}"
            size /= 1024.0
        return f"{size:.1f}PB"


class File(Node):
    def __init__(self, path: Path, name: str = None, size: int = 0):
        super().__init__(path, name)
        self.size = size

    @property
    def extension(self) -> str:
        """Get file extension"""
        return self.path.suffix

    def __str__(self):
        return f"File: {self.name} ({self._format_size(self.size)})"

    def _format_size(self, size: int) -> str:
        """Format size in human-readable format"""
        for unit in ['B', 'KB', 'MB', 'GB', 'TB']:
            if size < 1024.0:
                return f"{size:.1f}{unit}"
            size /= 1024.0
        return f"{size:.1f}PB"


def build_tree(root_path: Path, max_depth: int = None) -> Tree:
    """Build tree with optional depth limit"""
    if not root_path.exists():
        raise FileNotFoundError(f"Path does not exist: {root_path}")

    def _build_node(path: Path, current_depth: int = 0) -> Union[File, Dir]:
        if path.is_file():
            try:
                return File(path, size=path.stat().st_size)
            except (OSError, PermissionError):
                return File(path, size=0)
        elif path.is_dir():
            dir_node = Dir(path)

            # Check depth limit
            if max_depth is not None and current_depth >= max_depth:
                return dir_node

            try:
                for child_path in sorted(path.iterdir()):
                    try:
                        child_node = _build_node(child_path, current_depth + 1)
                        dir_node.add_child(child_node)
                    except (OSError, PermissionError):
                        # Skip files/directories we can't access
                        continue
            except PermissionError:
                # Handle directories we can't access
                pass
            return dir_node

    root_node = _build_node(root_path)
    return Tree(root_node)


def print_tree(node: Node, indent: str = "", show_size: bool = True, show_permissions: bool = False) -> None:
    """Enhanced tree printing with options"""
    output = indent + str(node)

    if show_permissions:
        try:
            output += f" [{node.permissions}]"
        except (OSError, PermissionError):
            output += " [???]"

    print(output)

    if hasattr(node, 'children'):
        for child in node.children:
            print_tree(child, indent + "  ", show_size, show_permissions)


def print_tree_statistics(tree: Tree) -> None:
    """Print statistics about the tree"""
    if isinstance(tree.root, Dir):
        print(f"Total files: {tree.root.file_count}")
        print(f"Total directories: {tree.root.dir_count}")
        print(f"Total size: {tree.root._format_size(tree.root.total_size)}")

        # Find largest files
        largest_files = tree.find_largest_files(5)
        if largest_files:
            print("\nLargest files:")
            for file in largest_files:
                print(f"  {file.name}: {file._format_size(file.size)}")


if __name__ == "__main__":
    root_path = Path(".")

    # Build tree with depth limit to avoid going too deep
    tree = build_tree(root_path, max_depth=3)

    # Print tree with enhanced options
    print("File Tree:")
    print_tree(tree.root, show_permissions=True)

    print("\n" + "="*50)
    print_tree_statistics(tree)

    # Example of finding specific files
    python_files = tree.find_files_by_extension('.py')
    if python_files:
        print(f"\nFound {len(python_files)} Python files:")
        for file in python_files[:5]:  # Show first 5
            print(f"  {file.name}")
