import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Checkbox, TableSortLabel } from '@material-ui/core';

interface Props {
  setSelectedProject: React.Dispatch<React.SetStateAction<number | undefined>>;
  setIsDrawerOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const useStyles = makeStyles({
  table: {
    minWidth: 650,
    cursor: 'pointer'
  }
});

const projects = ['Projekt 1', 'Projekt 2', 'Projekt 3', 'Projekt 4', 'Projekt 5'];

export default function ProjectTable({ setSelectedProject, setIsDrawerOpen }: Props) {
  const classes = useStyles();

  const handleCellClick = (selectedProject: number) => {
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
              <TableCell component="th" scope="row" onClick={() => handleCellClick(i)}>
                {project}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
