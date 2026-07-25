"""Tests for version_diff.py path bounding (REQ-SONAR-PG-07)."""

from __future__ import annotations

import pytest

from version_diff import PathOutsideRepoError, resolve_under_repo


def test_resolve_rejects_path_escape_outside_repo(tmp_path, monkeypatch):
    """Out-of-root paths like ../../etc/passwd must be rejected before open."""
    repo = tmp_path / "repo"
    repo.mkdir()
    monkeypatch.chdir(repo)

    with pytest.raises((PathOutsideRepoError, SystemExit)) as exc_info:
        resolve_under_repo("../../etc/passwd", repo_root=repo)

    if isinstance(exc_info.value, SystemExit):
        assert exc_info.value.code not in (0, None)


def test_resolve_accepts_path_under_repo(tmp_path, monkeypatch):
    repo = tmp_path / "repo"
    repo.mkdir()
    target = repo / "config" / "schema.json"
    target.parent.mkdir()
    target.write_text("{}")
    monkeypatch.chdir(repo)

    resolved = resolve_under_repo("config/schema.json", repo_root=repo)
    assert resolved == target.resolve()
    assert resolved.is_relative_to(repo.resolve())


def test_resolve_rejects_absolute_path_outside_repo(tmp_path):
    repo = tmp_path / "repo"
    repo.mkdir()
    outside = tmp_path / "outside.json"
    outside.write_text("{}")

    with pytest.raises((PathOutsideRepoError, SystemExit)) as exc_info:
        resolve_under_repo(str(outside), repo_root=repo)

    if isinstance(exc_info.value, SystemExit):
        assert exc_info.value.code not in (0, None)
