import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Checkbox, TableSortLabel } from '@material-ui/core';
import { ProjectDTO } from '../api/apis';

interface Props {
  projects: ProjectDTO[];
  setSelectedProject: React.Dispatch<React.SetStateAction<string | undefined>>;
  setIsDrawerOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const useStyles = makeStyles({
  table: {
    minWidth: 650,
    cursor: 'pointer'
  }
});

export default function ProjectTable({ setSelectedProject, setIsDrawerOpen, projects }: Props) {
  const classes = useStyles();

  const handleCellClick = (selectedProject: string) => {
    setIsDrawerOpen(true);
    setSelectedProject(selectedProject);
  };

  return (
    <TableContainer>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell padding="checkbox">
              <Checkbox />
            </TableCell>
            <TableCell>
              <TableSortLabel>Name</TableSortLabel>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {projects.map((project, i) => (
            <TableRow hover key={i}>
              <TableCell padding="checkbox">
                <Checkbox />
              </TableCell>
              <TableCell component="th" scope="row" onClick={() => handleCellClick(project.id)}>
                {project.name}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
